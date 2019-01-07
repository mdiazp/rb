import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Client } from '../../../../models/core';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-form-client',
  templateUrl: './form-client.component.html',
  styleUrls: ['./form-client.component.css']
})
export class FormClientComponent implements OnInit {

  @Input() initialValues = new Client(0, '', '', '', '', '', true);

  identificationControl: FormControl;
  nameControl: FormControl;
  addressControl: FormControl;
  phonesControl: FormControl;
  descriptionControl: FormControl;
  activedControl: FormControl;
  form: FormGroup;

  constructor() {}

  ngOnInit() {
    this.initForm();
  }

  initForm(): void {
    this.nameControl = new FormControl(
      this.initialValues.Name,
      [Validators.maxLength(100), Validators.required],
    );
    this.addressControl = new FormControl(
      this.initialValues.Address,
      [Validators.maxLength(1024), Validators.required],
    );
    this.identificationControl = new FormControl(
      this.initialValues.Identification,
      [Validators.maxLength(100)]
    );
    this.phonesControl = new FormControl(
      this.initialValues.Phones,
      [Validators.maxLength(100)],
    );
    this.descriptionControl = new FormControl(
      this.initialValues.Description,
      [Validators.maxLength(1024)],
    );
    this.activedControl = new FormControl(this.initialValues.Actived, Validators.required);
    this.form = new FormGroup({
      'name': this.nameControl,
      'address': this.addressControl,
      'identification': this.identificationControl,
      'phones': this.phonesControl,
      'description': this.descriptionControl,
      'actived': this.activedControl,
    });
  }

  public Valid(): boolean {
    return this.form.valid;
  }

  public ResetValues(): void {
    this.nameControl.setValue(this.initialValues.Name);
    this.addressControl.setValue(this.initialValues.Address);
    this.phonesControl.setValue(this.initialValues.Phones);
    this.identificationControl.setValue(this.initialValues.Identification);
    this.descriptionControl.setValue(this.initialValues.Description);
    this.activedControl.setValue(this.initialValues.Actived);
  }

  public GetClient(): Client {
    return new Client(
      this.initialValues.ID,
      this.identificationControl.value,
      this.nameControl.value,
      this.addressControl.value,
      this.phonesControl.value,
      this.descriptionControl.value,
      this.activedControl.value,
    );
  }
}
