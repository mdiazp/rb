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

@NgModule({
  imports: [
    CommonModule,
    DemoMaterialModule,
    SharedModule,
    MyCommonModule,

    RouterModule.forChild(DiscReservationsRoutes),
  ],
  declarations: [
    PDiscReservationsAllComponent,
    DiscReservationsDashboardComponent,
    DiscReservationsHistoryComponent,
    DiscReservationsCalendarComponent,
  ]
})
export class DiscReservationsModule { }
