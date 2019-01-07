import { Component, OnInit } from '@angular/core';
import { Disk } from '../../../models/core';
import { BehaviorSubject } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { APIDiskService, ErrorHandlerService } from '../../../services/core';

@Component({
  selector: 'app-disk-one',
  templateUrl: './disk-one.component.html',
  styleUrls: ['./disk-one.component.css']
})
export class DiskOneComponent implements OnInit {

  diskID: number;
  disk: Disk;

  private loadingSubject = new BehaviorSubject<boolean>(true);
  private loading$ = this.loadingSubject.asObservable();


  constructor(private router: Router,
              private route: ActivatedRoute,
              private api: APIDiskService,
              private eh: ErrorHandlerService) {
    this.route.params.subscribe(
      params => {
        this.diskID = params.id;
        this.loadDisk();
      }
    );
  }

  ngOnInit() {
    this.loadDisk();
  }

  refresh(): void {
    this.router.navigate(['/', 'discs', 'showone', this.disk.ID]);
  }

  loadDisk(): void {
    this.loadingSubject.next(true);
    this.api.GetDisk(this.diskID).subscribe(
      (disk) => {
        this.disk = disk;
        this.loadingSubject.next(false);
      },
      (e) => {
        this.router.navigate(['/', 'discs', 'all']);
        this.eh.HandleError(e);
      }
    );
  }
}
