import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Disk } from '../../../../models/core';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-form-disk',
  templateUrl: './form-disk.component.html',
  styleUrls: ['./form-disk.component.css']
})
export class FormDiskComponent implements OnInit {

  @Input() initialValues = new Disk(0, '', '', 1024, 'BIG', true);

  nameControl: FormControl;
  serialNumberControl: FormControl;
  capacityControl: FormControl;
  categoryControl: FormControl;
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
    this.serialNumberControl = new FormControl(
      this.initialValues.SerialNumber,
      [Validators.maxLength(255), Validators.required],
    );
    this.capacityControl = new FormControl(
      this.initialValues.Capacity,
      [Validators.maxLength(100), Validators.min(1), Validators.required]
    );
    this.categoryControl = new FormControl(
      this.initialValues.Category,
      [Validators.required],
    );
    this.activedControl = new FormControl(this.initialValues.Actived, Validators.required);
    this.form = new FormGroup({
      'name': this.nameControl,
      'serialNumber': this.serialNumberControl,
      'capacity': this.capacityControl,
      'category': this.categoryControl,
      'actived': this.activedControl,
    });
  }

  public Valid(): boolean {
    return this.form.valid;
  }

  public ResetValues(): void {
    this.nameControl.setValue(this.initialValues.Name);
    this.serialNumberControl.setValue(this.initialValues.SerialNumber);
    this.capacityControl.setValue(this.initialValues.Capacity);
    this.categoryControl.setValue(this.initialValues.Category);
    this.activedControl.setValue(this.initialValues.Actived);
  }

  public GetDisk(): Disk {
    return new Disk(
      this.initialValues.ID,
      this.nameControl.value,
      this.serialNumberControl.value,
      this.capacityControl.value,
      this.categoryControl.value,
      this.activedControl.value,
    );
  }
}
