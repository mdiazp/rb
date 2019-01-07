import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {Disk} from '../../models/core';
import {APIDiskService, DiskFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class DisksDataSource implements DataSource<Disk> {

    private disksSubject = new BehaviorSubject<Disk[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIDiskService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: DiskFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadDisks(filter);
      }
    }

    private loadCount(filter?: DiskFilter) {
      this.api.GetDisksCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadDisks(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.disksSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadDisks(filter?: DiskFilter) {
      this.api.GetDisks(filter).subscribe(
        disks => this.disksSubject.next(disks),
        (e) => {
          this.disksSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<Disk[]> {
        return this.disksSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.disksSubject.complete();
        this.loadingSubject.complete();
    }

}
