import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DiscReservationsAllComponent } from './disc-reservations-all/disc-reservations-all.component';
import { RouterModule } from '@angular/router';
import { DiscReservationsRoutes } from './disc-reservations.routing';
import { DemoMaterialModule } from '../../demo-material-module';
import { DiscReservationsDashboardComponent } from './disc-reservations-dashboard/disc-reservations-dashboard.component';
import { DiscReservationsHistoryComponent } from './disc-reservations-history/disc-reservations-history.component';
import { DiscReservationsCalendarComponent } from './disc-reservations-calendar/disc-reservations-calendar.component';

@NgModule({
  imports: [
    CommonModule,
    DemoMaterialModule,

    RouterModule.forChild(DiscReservationsRoutes),
  ],
  declarations: [
    DiscReservationsAllComponent,
    DiscReservationsDashboardComponent,
    DiscReservationsHistoryComponent,
    DiscReservationsCalendarComponent,
  ]
})
export class DiscReservationsModule { }
