import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {PDiscReservation} from '../../models/core';
import {APIPDiscReservationService, PDiscReservationFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class PDiscReservationsDataSource implements DataSource<PDiscReservation> {

    private pdiscReservationsSubject = new BehaviorSubject<PDiscReservation[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIPDiscReservationService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: PDiscReservationFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadPDiscReservations(filter);
      }
    }

    private loadCount(filter?: PDiscReservationFilter) {
      this.api.GetPDiscReservationsCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadPDiscReservations(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.pdiscReservationsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadPDiscReservations(filter?: PDiscReservationFilter) {
      this.api.GetPDiscReservations(filter).subscribe(
        PDiscReservations => this.pdiscReservationsSubject.next(PDiscReservations),
        (e) => {
          this.pdiscReservationsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<PDiscReservation[]> {
        return this.pdiscReservationsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.pdiscReservationsSubject.complete();
        this.loadingSubject.complete();
    }

}
