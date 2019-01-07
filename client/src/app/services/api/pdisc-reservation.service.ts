import { Injectable } from '@angular/core';
import {
  Http,
  URLSearchParams,
  Response,
  RequestOptions,
  RequestOptionsArgs,
  Headers,
} from '@angular/http';
import { BehaviorSubject, Observable, Operator } from 'rxjs';
import { map } from 'rxjs/operators';

import {
  PDiscReservation,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';

@Injectable()
export class APIPDiscReservationService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public PostPDiscReservation(pdiscReservation: PDiscReservation): Observable<PDiscReservation> {
    return this.post('/pdiskreservation', pdiscReservation);
  }

  public GetPDiscReservation(id: number): Observable<PDiscReservation> {
    return this.get(`/pdiskreservation/${id}`);
  }

  public UpdatePDiscReservation(pdiscReservation: PDiscReservation): Observable<PDiscReservation> {
    return this.patch(`/pdiskreservation/${pdiscReservation.ID}`, pdiscReservation);
  }

  public DeletePDiscReservation(pdiscReservationID: number): Observable<Response> {
    return this.delete(`/pdiskreservation/${pdiscReservationID}`);
  }

  public GetPDiscReservations(filter?: PDiscReservationFilter): Observable<PDiscReservation[]> {
    if ( filter && filter !== null ) {
      return this.get('/pdiskreservations', { params: filter.GetUSP() });
    } else {
      return this.get('/pdiskreservations');
    }
  }

  public GetPDiscReservationsCount(filter?: PDiscReservationFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/pdiskreservationscount', { params: filter.GetUSP() });
    } else {
      return this.get('/pdiskreservationscount');
    }
  }
}

export class PDiscReservationFilter {
  constructor(
    public ClientID: number,
    public ActivedClient: boolean,
    public ActivedInitialTime: Date,
    public ActivedFinishTime: Date,
    public DiskCategoryRequest: string,
    public TurnWeekDay: string,
    public TurnNum: number,
    public paginator: Paginator,
    public orderby: OrderBy
    ) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.ClientID) && this.ClientID !== 0 ) {
      usp.append('clientID', this.ClientID.toString());
    }
    if ( !isNullOrUndefined(this.ActivedClient) ) {
      usp.append('activedClient', this.ActivedClient.toString());
    }
    if ( !isNullOrUndefined(this.ActivedInitialTime) ) {
      usp.append('activedInitialTime', this.ActivedInitialTime.toString());
    }
    if ( !isNullOrUndefined(this.ActivedFinishTime) ) {
      usp.append('activedFinishTime', this.ActivedFinishTime.toString());
    }
    if ( !isNullOrUndefined(this.DiskCategoryRequest) && this.DiskCategoryRequest !== '' ) {
      usp.append('diskCategoryRequest', this.DiskCategoryRequest.toString());
    }
    if ( !isNullOrUndefined(this.TurnWeekDay) && this.TurnWeekDay !== '' ) {
      usp.append('turnWeekDay', this.TurnWeekDay.toString());
    }
    if ( !isNullOrUndefined(this.TurnNum) && this.TurnNum !== 0 ) {
      usp.append('turnNum', this.TurnNum.toString());
    }
    if ( !isNullOrUndefined(this.paginator) ) {
      usp.appendAll(this.paginator.GetUSP());
    }
    if ( !isNullOrUndefined(this.orderby) ) {
      usp.appendAll(this.orderby.GetUSP());
    }
    return usp;
  }
}
