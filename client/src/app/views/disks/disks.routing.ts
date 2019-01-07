import { Routes } from '@angular/router';
import { DiskAllComponent } from './disk-all/disk-all.component';
import { DiskOneComponent } from './disk-one/disk-one.component';

export const DisksRoutes: Routes = [
  {
    path: '',
    redirectTo: 'all',
    pathMatch: 'full',
  },
  {
    path: 'all',
    component: DiskAllComponent,
  },
  {
    path: 'showone/:id',
    component: DiskOneComponent,
  }
];
