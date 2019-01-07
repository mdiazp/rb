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
  Client,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';

@Injectable()
export class APIClientService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public PostClient(client: Client): Observable<Client> {
    return this.post('/client', client);
  }

  public GetClient(id: number): Observable<Client> {
    return this.get(`/client/${id}`);
  }

  public UpdateClient(client: Client): Observable<Client> {
    return this.patch(`/client/${client.ID}`, client);
  }

  public DeleteClient(clientID: number): Observable<Response> {
    return this.delete(`/client/${clientID}`);
  }

  public GetClients(filter?: ClientFilter): Observable<Client[]> {
    if ( filter && filter !== null ) {
      return this.get('/clients', { params: filter.GetUSP() });
    } else {
      return this.get('/clients');
    }
  }

  public GetClientsCount(filter?: ClientFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/clientscount', { params: filter.GetUSP() });
    } else {
      return this.get('/clientscount');
    }
  }
}

export class ClientFilter {
  constructor(
    public IdentificationPrefix: string,
    public NameSubstr: string,
    public AddressSubstr: string,
    public PhonesSubstr: string,
    public DescriptionSubstr: string,
    public Actived: boolean,
    public paginator: Paginator,
    public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.IdentificationPrefix) && this.IdentificationPrefix !== '' ) {
      usp.append('identificationPrefix', this.IdentificationPrefix.toString());
    }
    if ( !isNullOrUndefined(this.NameSubstr) && this.NameSubstr !== '' ) {
      usp.append('nameSubstr', this.NameSubstr.toString());
    }
    if ( !isNullOrUndefined(this.AddressSubstr) && this.AddressSubstr !== '' ) {
      usp.append('addressSubstr', this.AddressSubstr.toString());
    }
    if ( !isNullOrUndefined(this.PhonesSubstr) && this.PhonesSubstr !== '' ) {
      usp.append('phonesSubstr', this.PhonesSubstr.toString());
    }
    if ( !isNullOrUndefined(this.DescriptionSubstr) && this.DescriptionSubstr !== '' ) {
      usp.append('descriptionSubstr', this.DescriptionSubstr.toString());
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
