import { Component, OnInit, Input } from '@angular/core';
import { Disk } from '../../../../models/core';

@Component({
  selector: 'app-disk-info',
  templateUrl: './disk-info.component.html',
  styleUrls: ['./disk-info.component.css']
})
export class DiskInfoComponent implements OnInit {

  @Input() disk: Disk;

  constructor() { }

  ngOnInit() {
  }

}
