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
  PDiscReservation, Util, PDRTurnCalendar,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';
import { discardPeriodicTasks } from '@angular/core/testing';

@Injectable()
export class APIPDiscReservationService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService,
              protected util: Util) {
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
      return this.get('/pdiskreservations', { params: filter.GetUSP() }).pipe(
        map(res => this.parseList(res)),
      );
    } else {
      return this.get('/pdiskreservations').pipe(
        map(res => this.parseList(res)),
      );
    }
  }

  public GetCalendar(weekDay: string, turnNum: number): Observable<PDRTurnCalendar> {
    return this.get(`/pdisk-reservation/calendar/${weekDay}/${turnNum}`).pipe(
      map(res => {
        for (let i = 0; i < res.States.length; i++) {
          res.States[i].Date = this.util.NewDate(res.States[i].Date);
          res.States[i].PDRs = this.parseList(res.States[i].PDRs);
        }
        return res;
      })
    );
  }

  parseList(list: any[]): PDiscReservation[] {
    for (let i = 0; i < list.length; i++ ) {
      list[i].InitialTime = this.util.NewDate(list[i].InitialTime);
      list[i].FinishTime = this.util.NewDate(list[i].FinishTime);
      list[i].TurnWeekDay = this.util.GetWeekDayDisplayValue(list[i].TurnWeekDay);
      list[i].DiskCategoryRequest = this.util.GetDiscCategoryDisplayValue(list[i].DiskCategoryRequest);
    }
    return list;
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
  private util = new Util();

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
      usp.append(
        'activedInitialTime',
        this.util.FormatDateToSendToApi(this.ActivedInitialTime)
      );
    }
    if ( !isNullOrUndefined(this.ActivedFinishTime) ) {
      usp.append(
        'activedFinishTime',
        this.util.FormatDateToSendToApi(this.ActivedFinishTime)
      );
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
