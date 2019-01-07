import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {Client} from '../../models/core';
import {APIClientService, ClientFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class ClientsDataSource implements DataSource<Client> {

    private clientsSubject = new BehaviorSubject<Client[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIClientService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: ClientFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadClients(filter);
      }
    }

    private loadCount(filter?: ClientFilter) {
      this.api.GetClientsCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadClients(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.clientsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadClients(filter?: ClientFilter) {
      this.api.GetClients(filter).subscribe(
        clients => this.clientsSubject.next(clients),
        (e) => {
          this.clientsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<Client[]> {
        return this.clientsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.clientsSubject.complete();
        this.loadingSubject.complete();
    }

}
