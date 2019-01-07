import { Component, OnInit } from '@angular/core';
import { Client } from '../../../models/core';
import { BehaviorSubject } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { APIClientService, ErrorHandlerService } from '../../../services/core';

@Component({
  selector: 'app-client-one',
  templateUrl: './client-one.component.html',
  styleUrls: ['./client-one.component.scss']
})
export class ClientOneComponent implements OnInit {

  clientID: number;
  client: Client;

  private loadingSubject = new BehaviorSubject<boolean>(true);
  private loading$ = this.loadingSubject.asObservable();


  constructor(private router: Router,
              private route: ActivatedRoute,
              private api: APIClientService,
              private eh: ErrorHandlerService) {
    this.route.params.subscribe(
      params => {
        this.clientID = params.id;
        this.loadClient();
      }
    );
  }

  ngOnInit() {
    this.loadClient();
  }

  refresh(): void {
    this.router.navigate(['/', 'clients', 'showone', this.client.ID]);
  }

  loadClient(): void {
    this.loadingSubject.next(true);
    this.api.GetClient(this.clientID).subscribe(
      (client) => {
        this.client = client;
        this.loadingSubject.next(false);
      },
      (e) => {
        this.router.navigate(['/', 'clients', 'all']);
        this.eh.HandleError(e);
      }
    );
  }
}
