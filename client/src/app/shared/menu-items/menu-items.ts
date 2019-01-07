import { Injectable } from '@angular/core';
import { SessionService } from '../../services/core';
import { RolAdmin } from '../../models/core';

export interface Menu {
  state: string[];
  name: string;
  type: string;
  icon: string;
}

const MENUITEMS_ROLADMIN = [
  { state: ['/', 'home'], name: 'Home', type: 'link', icon: 'home' },
  { state: ['/', 'discs'], name: 'Discos', type: 'link', icon: 'album' },
  { state: ['/', 'clients'], name: 'Clientes', type: 'link', icon: 'people' },
  { state: ['/', 'discreservations'], name: 'Reservacion de Disco', type: 'link', icon: 'people' },
  { state: ['/', 'disccopies'], name: 'Copia de Disco', type: 'link', icon: 'people' },
];

@Injectable()
export class MenuItems {
  menu: Menu[] = [];

  constructor(private session: SessionService) {
    this.menu.push(
      { state: ['/', 'home'], name: 'Home', type: 'link', icon: 'home' },
      { state: ['/', 'discs'], name: 'Discos', type: 'link', icon: 'album' },
      { state: ['/', 'clients'], name: 'Clientes', type: 'link', icon: 'people' },
      { state: ['/', 'discreservations'], name: 'Reservacion de Disco', type: 'link', icon: 'date_range' },
      { state: ['/', 'disccopies'], name: 'Copia de Disco', type: 'link', icon: 'event' },
    );
  }

  getMenuitem(): Menu[] {
    if ( this.session.GetUser().Rol === RolAdmin ) {
      return MENUITEMS_ROLADMIN;
    } else {
      return this.menu;
    }
  }
}
