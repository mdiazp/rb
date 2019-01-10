import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import {
  ErrorHandlerService,
  APIClientService,
  ClientFilter,
  Paginator
} from '../../../services/core';
import { Client } from '../../../models/core';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { MatAutocompleteSelectedEvent } from '@angular/material';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-client-selector',
  templateUrl: './client-selector.component.html',
  styleUrls: ['./client-selector.component.css']
})
export class ClientSelectorComponent implements OnInit {

  form: FormGroup;
  clientNameControl: FormControl;

  clients: Client[] = [];
  autoClientSelection = new Client(0, null, null, null, null, null, null);

  @Output() selectionChanges = new EventEmitter<Client>();

  constructor(private api: APIClientService,
              private eh: ErrorHandlerService) { }

  ngOnInit() {
    this.clientNameControl = new FormControl('');
    this.form = new FormGroup({
      'clientName': this.clientNameControl,
    });

    this.clientNameControl.valueChanges
      .pipe(
        debounceTime(150),
        distinctUntilChanged(),
      )
      .subscribe(
        (value) => {
          if (value !== '') {
            this.loadClients(value);
          } else {
            this.clients = [];
          }

          this.selectionChanges.emit(this.autoClientSelection);
        }
      );
  }

  ValidSelection(): boolean {
    return (!isNullOrUndefined(this.autoClientSelection.Name) &&
            this.autoClientSelection.Name !== '' &&
            this.autoClientSelection === this.clientNameControl.value);
  }

  Clear(): void {
    this.clientNameControl.setValue('');
  }

  loadClients(value: string): void {
    this.api.GetClients(
      new ClientFilter(null, value, null, null, null, null,
      new Paginator(0, 10), null)
    ).subscribe(
      (data) => {
        this.clients = data;
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  onSelectClient(ev: MatAutocompleteSelectedEvent) {
    this.autoClientSelection = ev.option.value;
  }

  displayClient(client?: Client): string | undefined {
    return client ? client.Name : undefined;
  }
}
