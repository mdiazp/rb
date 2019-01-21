import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PDiscReservationsAllComponent } from './pdisc-reservations-all/pdisc-reservations-all.component';
import { RouterModule } from '@angular/router';

import { MyCommonModule } from '../my-common/my-common.module';

import { DiscReservationsRoutes } from './disc-reservations.routing';
import { DemoMaterialModule } from '../../demo-material-module';
import { DiscReservationsDashboardComponent } from './disc-reservations-dashboard/disc-reservations-dashboard.component';
import { DiscReservationsHistoryComponent } from './disc-reservations-history/disc-reservations-history.component';
import { DiscReservationsCalendarComponent } from './disc-reservations-calendar/disc-reservations-calendar.component';
import { SharedModule } from '../../shared/shared.module';
import { NewDiscReservationDialogComponent } from './common/new-disc-reservation-dialog/new-disc-reservation-dialog.component';
import {
  ViewOneDiscReservationDialogComponent,
} from './common/view-one-disc-reservation-dialog/view-one-disc-reservation-dialog.component';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  imports: [
    CommonModule,
    DemoMaterialModule,
    SharedModule,
    MyCommonModule,
    ReactiveFormsModule,

    RouterModule.forChild(DiscReservationsRoutes),
  ],
  declarations: [
    PDiscReservationsAllComponent,
    DiscReservationsDashboardComponent,
    DiscReservationsHistoryComponent,
    DiscReservationsCalendarComponent,
    NewDiscReservationDialogComponent,
    ViewOneDiscReservationDialogComponent,
  ],
  entryComponents: [
    NewDiscReservationDialogComponent,
    ViewOneDiscReservationDialogComponent,
  ],
})
export class DiscReservationsModule { }
