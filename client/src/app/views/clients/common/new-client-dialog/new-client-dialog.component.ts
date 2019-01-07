import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Client } from '../../../../models/core';
import { FormClientComponent } from '../form-client/form-client.component';

@Component({
  selector: 'app-new-client-dialog',
  templateUrl: './new-client-dialog.component.html',
  styleUrls: ['./new-client-dialog.component.scss']
})
export class NewClientDialogComponent implements OnInit {
  @ViewChild(FormClientComponent) form: FormClientComponent;

  constructor() { }

  ngOnInit() {
  }

  private getClient(): Client {
    return this.form.GetClient();
  }
}
