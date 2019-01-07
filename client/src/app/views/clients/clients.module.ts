import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { ClientOneComponent } from './client-one/client-one.component';
import { ClientAllComponent } from './client-all/client-all.component';
import { ClientsRoutes } from './clients.routing';

import { DemoMaterialModule } from '../../demo-material-module';
import { NewClientDialogComponent } from './common/new-client-dialog/new-client-dialog.component';
import { FormClientComponent } from './common/form-client/form-client.component';
import { ClientSettingsComponent } from './common/client-settings/client-settings.component';
import { ClientInfoComponent } from './common/client-info/client-info.component';
import { ClientServicesComponent } from './common/client-services/client-services.component';

@NgModule({
  imports: [
    CommonModule,
    DemoMaterialModule,
    ReactiveFormsModule,
    // FormsModule,
    RouterModule.forChild(ClientsRoutes),
  ],
  entryComponents: [
    NewClientDialogComponent,
  ],
  declarations: [
    ClientOneComponent,
    ClientAllComponent,
    NewClientDialogComponent,
    FormClientComponent,
    ClientSettingsComponent,
    ClientInfoComponent,
    ClientServicesComponent
  ]
})
export class ClientsModule { }
