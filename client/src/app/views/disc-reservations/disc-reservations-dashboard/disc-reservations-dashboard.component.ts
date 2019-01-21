import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material';

import {
  NewDiscReservationDialogComponent,
} from '../common/new-disc-reservation-dialog/new-disc-reservation-dialog.component';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-disc-reservations-dashboard',
  templateUrl: './disc-reservations-dashboard.component.html',
  styleUrls: ['./disc-reservations-dashboard.component.css']
})
export class DiscReservationsDashboardComponent implements OnInit {

  links = [
    {
      Path: ['all'],
      Text: 'Lista',
      Icon: 'book',
    },
    {
      Path: ['calendar'],
      Text: 'Calendario',
      Icon: 'event_note',
    },
    {
      Path: ['history'],
      Text: 'Historial',
      Icon: 'history',
    },
  ];
  activeLink = this.links[0];

  constructor(private dialog: MatDialog) { }

  ngOnInit() {
  }

  onNew(): void {
    const ref = this.dialog.open(
      NewDiscReservationDialogComponent,
      {
        data: {
          routeToGetBack: ['dashboard'],
        }
      }
    );
  }

}
