import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

import { DemoMaterialModule } from '../../demo-material-module';
import { ClientSelectorComponent } from './client-selector/client-selector.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    DemoMaterialModule,
  ],
  exports: [
    ClientSelectorComponent,
  ],
  declarations: [
    ClientSelectorComponent,
  ],
})
export class MyCommonModule { }
