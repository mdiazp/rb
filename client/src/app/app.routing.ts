import { Routes } from '@angular/router';

import { FullComponent } from './layouts/full/full.component';
import { AuthGuard, RolAdminGuard } from './guards/core';
// import { LoginComponent } from './login/login.component';

export const AppRoutes: Routes = [
  {
    path: 'login',
    loadChildren: './views/login/login.module#LoginModule'
  },
  {
    path: '',
    component: FullComponent,
    children: [
      {
        path: '',
        redirectTo: '/home',
        pathMatch: 'full'
      },
      {
        path: 'home',
        loadChildren: './views/home/home.module#HomeModule'
      },
      {
        path: 'discs',
        loadChildren: './views/disks/disks.module#DisksModule'
      },
      {
        path: 'clients',
        loadChildren: './views/clients/clients.module#ClientsModule'
      },
      {
        path: 'discreservations',
        loadChildren: './views/disc-reservations/disc-reservations.module#DiscReservationsModule'
      },
    ]
  }
];
