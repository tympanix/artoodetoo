import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styles: []
})
export class DashboardComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  public tasks = [
    {
      taskName: "Facebook Tag to Google Drive Save",
      event: "Facebook Image Tag",
      converters: [
        {
          converterName: "Watermark"
        },
        {
          converterName: "Zip compression"
        },
        {
          converterName: "Zip compression"
        },
        {
          converterName: "Zip compression"
        },
        {
          converterName: "Zip compression"
        }
      ],
      action: "Google Drive Save"
    },
    {
      taskName: "Facebook Tag to Google Drive Save",
      event: "Facebook Image Tag",
      converters: [
        {
          converterName: "Watermark"
        },
        {
          converterName: "Zip compression"
        }
      ],
      action: "Google Drive Save"
    },
    {
      taskName: "Facebook Tag to Google Drive Save",
      event: "Facebook Image Tag",
      converters: [
        {
          converterName: "Watermark"
        },
        {
          converterName: "Zip compression"
        }
      ],
      action: "Google Drive Save"
    }
  ]

}
