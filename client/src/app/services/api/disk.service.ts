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
  Disk,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';

@Injectable()
export class APIDiskService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public PostDisk(disk: Disk): Observable<Disk> {
    return this.post('/disk', disk);
  }

  public GetDisk(id: number): Observable<Disk> {
    return this.get(`/disk/${id}`);
  }

  public UpdateDisk(disk: Disk): Observable<Disk> {
    return this.patch(`/disk/${disk.ID}`, disk);
  }

  public DeleteDisk(diskID: number): Observable<Response> {
    return this.delete(`/disk/${diskID}`);
  }

  public GetDisks(filter?: DiskFilter): Observable<Disk[]> {
    if ( filter && filter !== null ) {
      return this.get('/disks', { params: filter.GetUSP() }).pipe(
        map(res => this.parseList(res))
      );
    } else {
      return this.get('/disks').pipe(
        map(res => this.parseList(res))
      );
    }
  }

  parseList(discs: Disk[]): Disk[] {
    for (let i = 0; i < discs.length; i++) {
      discs[i] = this.parseDisk(discs[i]);
    }
    return discs;
  }

  parseDisk(d: Disk): Disk {
    d.Category = this.util.GetDiscCategoryDisplayValue(d.Category);
    return d;
  }

  public GetDisksCount(filter?: DiskFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/diskscount', { params: filter.GetUSP() });
    } else {
      return this.get('/diskscount');
    }
  }
}

export class DiskFilter {
  constructor(
    public Capacity: number,
    public Category: string,
    public Actived: boolean,
    public paginator: Paginator,
    public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.Capacity) && this.Capacity !== 0 ) {
      usp.append('capacity', this.Capacity.toString());
    }
    if ( !isNullOrUndefined(this.Category) && this.Category !== '' ) {
      usp.append('category', this.Category.toString());
    }
    if ( !isNullOrUndefined(this.Actived) ) {
      usp.append('actived', this.Actived.toString());
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
