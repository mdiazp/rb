import { Component, ViewChild, OnInit, AfterViewInit, ElementRef, Input } from '@angular/core';
import { MatPaginator, MatSort, MatSelect, MatInput } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent } from 'rxjs';
import { MatDialog } from '@angular/material';

import { CheckDeleteDialogComponent } from '../../../dialogs/core';

import {
  APIDiskService,
  DiskFilter,
  DisksDataSource,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  OrderBy,
} from '../../../services/core';
import { Disk } from '../../../models/core';
import { isNullOrUndefined } from 'util';
import { NewDiskDialogComponent } from '../common/new-disk-dialog/new-disk-dialog.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-disk-all',
  templateUrl: './disk-all.component.html',
  styleUrls: ['./disk-all.component.scss'],
})
export class DiskAllComponent implements OnInit, AfterViewInit {

  dataSource: DisksDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['name', 'capacity', 'operations'];


  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;

  @ViewChild('capacityFilter') capacityFilter: ElementRef;
  @ViewChild('categoryFilter') categoryFilter: MatSelect;
  @ViewChild('activedFilter') activedFilter: MatSelect;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIDiskService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private router: Router,
              private dialog: MatDialog) {}

  ngOnInit() {
      this.dataSource = new DisksDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new DiskFilter(
          null, null, true,
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
    this.addKeyupEventToFilter(this.capacityFilter);

    this.categoryFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        this.load(true);
      },
    );

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
    const ref = this.dialog.open(NewDiskDialogComponent);

    ref.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) ) {
        this.api.PostDisk(result).subscribe(
          (disk) => {
            console.log('El nuevo disco fue creado exitosamente');
            this.fh.ShowFeedback('El nuevo disco fue creado exitosamente');
            this.router.navigate(['/', 'discs', 'showone', disk.ID]);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  onDelete(o: Disk): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Si elimina el disco ${o.Name} se
        perderan todos los datos relacionados con el mismo
        (reservaciones del disco, etc).
        Esta seguro de eliminar el disco?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteDisk(o.ID).subscribe(
          (_) => {
            console.log('El disco fue eliminado exitosamente.');
            this.fh.ShowFeedback('El disco fue eliminado exitosamente.');
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
      new DiskFilter(
        Number(this.capacityFilter.nativeElement.value),
        (this.categoryFilter.value === 'all' ? null : this.categoryFilter.value),
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
