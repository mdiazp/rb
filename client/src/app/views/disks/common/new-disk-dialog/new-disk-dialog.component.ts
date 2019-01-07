import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Disk } from '../../../../models/core';
import { FormDiskComponent } from '../form-disk/form-disk.component';

@Component({
  selector: 'app-new-disk-dialog',
  templateUrl: './new-disk-dialog.component.html',
  styleUrls: ['./new-disk-dialog.component.scss']
})
export class NewDiskDialogComponent implements OnInit {
  @ViewChild(FormDiskComponent) form: FormDiskComponent;

  constructor() { }

  ngOnInit() {
  }

  private getDisk(): Disk {
    return this.form.GetDisk();
  }
}
