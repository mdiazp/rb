import { Routes } from '@angular/router';
import { ClientAllComponent } from './client-all/client-all.component';
import { ClientOneComponent } from './client-one/client-one.component';

export const ClientsRoutes: Routes = [
  {
    path: '',
    redirectTo: 'all',
    pathMatch: 'full',
  },
  {
    path: 'all',
    component: ClientAllComponent,
  },
  {
    path: 'showone/:id',
    component: ClientOneComponent,
  }
];
