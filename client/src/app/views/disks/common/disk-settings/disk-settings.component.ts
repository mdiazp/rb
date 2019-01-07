import { Component, OnInit, Input, Output, EventEmitter, ViewChild, OnDestroy } from '@angular/core';

import { FormDiskComponent } from '../form-disk/form-disk.component';
import { Disk } from '../../../../models/core';
import { APIDiskService, ErrorHandlerService, FeedbackHandlerService } from '../../../../services/core';


@Component({
  selector: 'app-disk-settings',
  templateUrl: './disk-settings.component.html',
  styleUrls: ['./disk-settings.component.scss']
})
export class DiskSettingsComponent implements OnInit, OnDestroy {

  @Input() disk: Disk;
  @Output() change = new EventEmitter<boolean>();
  @ViewChild(FormDiskComponent) form: FormDiskComponent;

  constructor(private api: APIDiskService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService) {}

  ngOnInit(): void {
  }

  onSave(): void {
    const g = this.form.GetDisk();

    this.api.UpdateDisk(g).subscribe(
      (disk) => {
        this.disk.Name = disk.Name;
        this.disk.SerialNumber = disk.SerialNumber;
        this.disk.Capacity = disk.Capacity;
        this.disk.Category = disk.Category;
        this.disk.Actived = disk.Actived;

        this.change.emit(true);

        this.fh.ShowFeedback('El disco fue actualizado exitosamente');
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  ngOnDestroy(): void {
  }
}
