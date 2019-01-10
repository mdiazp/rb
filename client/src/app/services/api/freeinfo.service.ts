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
  Credentials,
  Session,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';

@Injectable()
export class APIFreeInfoService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  GetServerTime(): Observable<Date> {
    return this.get('/server-time').pipe(
      map(res => new Date(res)),
    );
  }

  GetTurnNums(): Observable<number[]> {
    return this.get('/turn-nums');
  }

  GetDiscCategoriesInfo(): Observable<DiscCategoriesInfo> {
    return this.get('/disc-categories-info');
  }
}

export class DiscCategoriesInfo {
  constructor(public DiscCategories: string[],
              public DiscCategoryRequestNull: string) {}
}
