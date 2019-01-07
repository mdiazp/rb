import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { DemoMaterialModule } from '../../demo-material-module';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { SharedModule } from '../../shared/shared.module';

import { DiskAllComponent } from './disk-all/disk-all.component';
import { DiskOneComponent } from './disk-one/disk-one.component';

import { DisksRoutes } from './disks.routing';
import { NewDiskDialogComponent } from './common/new-disk-dialog/new-disk-dialog.component';
import { FormDiskComponent } from './common/form-disk/form-disk.component';
import { DiskInfoComponent } from './common/disk-info/disk-info.component';
import { DiskSettingsComponent } from './common/disk-settings/disk-settings.component';

@NgModule({
  imports: [
    CommonModule,
    SharedModule,
    DemoMaterialModule,
    ReactiveFormsModule,

    RouterModule.forChild(DisksRoutes),
  ],
  declarations: [
    DiskAllComponent,
    DiskOneComponent,
    NewDiskDialogComponent,
    FormDiskComponent,
    DiskInfoComponent,
    DiskSettingsComponent
  ],
  entryComponents: [
    NewDiskDialogComponent,
  ],
})
export class DisksModule { }
