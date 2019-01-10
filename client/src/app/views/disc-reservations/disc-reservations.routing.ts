import { Routes } from '@angular/router';
import { PDiscReservationsAllComponent } from './pdisc-reservations-all/pdisc-reservations-all.component';
import { DiscReservationsDashboardComponent } from './disc-reservations-dashboard/disc-reservations-dashboard.component';
import { DiscReservationsHistoryComponent } from './disc-reservations-history/disc-reservations-history.component';
import { DiscReservationsCalendarComponent } from './disc-reservations-calendar/disc-reservations-calendar.component';

export const DiscReservationsRoutes: Routes = [
  {
    path: '',
    redirectTo: 'dashboard',
    pathMatch: 'full',
  },
  {
    path: 'dashboard',
    component: DiscReservationsDashboardComponent,
    children: [
      {path: '', redirectTo: 'all'},
      {path: 'all', component: PDiscReservationsAllComponent},
      {path: 'calendar', component: DiscReservationsCalendarComponent},
      {path: 'history', component: DiscReservationsHistoryComponent},
    ]
  },
];
