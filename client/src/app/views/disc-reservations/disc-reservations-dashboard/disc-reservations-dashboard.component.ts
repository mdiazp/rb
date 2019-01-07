import { Component, OnInit } from '@angular/core';

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

  constructor() { }

  ngOnInit() {
  }

  onNew(): void {

  }
}
