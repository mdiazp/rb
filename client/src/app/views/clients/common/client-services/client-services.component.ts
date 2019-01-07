import { Component, OnInit, Input } from '@angular/core';
import { Client } from '../../../../models/core';

@Component({
  selector: 'app-client-services',
  templateUrl: './client-services.component.html',
  styleUrls: ['./client-services.component.css']
})
export class ClientServicesComponent implements OnInit {

  @Input() client: Client;

  constructor() { }

  ngOnInit() {
  }

}
