import { Component, ViewChild, OnInit, AfterViewInit, ElementRef } from '@angular/core';
import { MatPaginator, MatSort, MatSelect } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent } from 'rxjs';
import { MatDialog } from '@angular/material';

import { CheckDeleteDialogComponent } from '../../../dialogs/core';

import {
  APIClientService,
  ClientFilter,
  ClientsDataSource,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  OrderBy,
} from '../../../services/core';
import { Client } from '../../../models/core';
import { isNullOrUndefined } from 'util';
import { NewClientDialogComponent } from '../common/new-client-dialog/new-client-dialog.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-client-all',
  templateUrl: './client-all.component.html',
  styleUrls: ['./client-all.component.scss'],
})
export class ClientAllComponent implements OnInit, AfterViewInit {

  dataSource: ClientsDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['name', 'operations'];


  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;

  @ViewChild('identificationFilter') identificationFilter: ElementRef;
  @ViewChild('nameFilter') nameFilter: ElementRef;
  @ViewChild('addressFilter') addressFilter: ElementRef;
  @ViewChild('phoneFilter') phoneFilter: ElementRef;
  @ViewChild('descriptionFilter') descriptionFilter: ElementRef;
  @ViewChild(MatSelect) activedFilter: MatSelect;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIClientService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private router: Router,
              private dialog: MatDialog) {}

  ngOnInit() {
      this.dataSource = new ClientsDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new ClientFilter(
          null, null, null, null, null, true,
          new Paginator(
            0,
            this.initialPageSize,
          ),
          new OrderBy(
            'id',
            false,
          )
        ),
      );
  }

  ngAfterViewInit() {
    this.addKeyupEventToFilter(this.identificationFilter);
    this.addKeyupEventToFilter(this.nameFilter);
    this.addKeyupEventToFilter(this.addressFilter);
    this.addKeyupEventToFilter(this.phoneFilter);
    this.addKeyupEventToFilter(this.descriptionFilter);

    this.activedFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        this.load(true);
      },
    );

    this.sort.sortChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        this.load(false);
      }
    );

    this.paginator.page
      .pipe(
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  addKeyupEventToFilter(controlFilter: ElementRef<any>) {
    fromEvent(controlFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    this.load(true);
                })
            )
            .subscribe();
  }

  onNew(): void {
    console.log('onNew');
    const ref = this.dialog.open(NewClientDialogComponent);

    ref.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) ) {
        this.api.PostClient(result).subscribe(
          (client) => {
            console.log('El nuevo cliente fue creado exitosamente');
            this.fh.ShowFeedback('El nuevo cliente fue creado exitosamente');
            this.router.navigate(['/', 'clients', 'showone', client.ID]);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  onDelete(o: Client): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Si elimina el cliente ${o.Name} se
        perderan todos los datos relacionados con el mismo
        (pagos, servicios brindados, etc).
        Esta seguro de eliminar el cliente?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteClient(o.ID).subscribe(
          (_) => {
            console.log('El cliente fue eliminado exitosamente.');
            this.fh.ShowFeedback('El cliente fue eliminado exitosamente.');
            this.paginator.pageIndex = 0;
            // this.paginator.page.emit();
            this.load(true);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  load(loadCount: boolean) {
    this.dataSource.load(
      loadCount,
      new ClientFilter(
        this.identificationFilter.nativeElement.value,
        this.nameFilter.nativeElement.value,
        this.addressFilter.nativeElement.value,
        this.phoneFilter.nativeElement.value,
        this.descriptionFilter.nativeElement.value,
        (this.activedFilter.value === 'all' ? null : this.activedFilter.value),
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
        new OrderBy(
          this.sort.active,
          (this.sort.direction !== 'asc'),
        ),
      ),
    );
  }
}
