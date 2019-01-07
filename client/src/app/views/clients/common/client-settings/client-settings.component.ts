import { Component, OnInit, Input, Output, EventEmitter, ViewChild, OnDestroy } from '@angular/core';

import { FormClientComponent } from '../form-client/form-client.component';
import { Client } from '../../../../models/core';
import { APIClientService, ErrorHandlerService, FeedbackHandlerService } from '../../../../services/core';


@Component({
  selector: 'app-client-settings',
  templateUrl: './client-settings.component.html',
  styleUrls: ['./client-settings.component.scss']
})
export class ClientSettingsComponent implements OnInit, OnDestroy {

  @Input() client: Client;
  @Output() change = new EventEmitter<boolean>();
  @ViewChild(FormClientComponent) form: FormClientComponent;

  constructor(private api: APIClientService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService) {}

  ngOnInit(): void {
  }

  onSave(): void {
    const g = this.form.GetClient();

    this.api.UpdateClient(g).subscribe(
      (client) => {
        this.client.Name = client.Name;
        this.client.Identification = client.Identification;
        this.client.Address = client.Address;
        this.client.Phones = client.Phones;
        this.client.Description = client.Description;
        this.client.Actived = client.Actived;

        this.change.emit(true);

        this.fh.ShowFeedback('El cliente fue actualizado exitosamente');
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  ngOnDestroy(): void {
  }
}
