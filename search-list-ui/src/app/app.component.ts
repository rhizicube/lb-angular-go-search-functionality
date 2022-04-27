import { Component } from '@angular/core';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  constructor(private appService: AppService) {}
  title: string = '';
  inputValue: string = '';
  isError: boolean = false;
  isLoading: boolean = false;
  isDataAvailable: boolean = false;
  t: Array<any> = [1, 2, 3];
  movies: Array<any> = [];
  filteredLists: Array<any> = [];
  frontendHandle: boolean = true;
  storedData: Array<any> = [];

  ngOnInit(): void {
    this.isLoading = true;
    this.appService.getUsers().subscribe((users: any) => {
      if ('data' in users) {
        this.isLoading = false;
        this.isDataAvailable = true;
        this.movies = users.data;
        this.filteredLists = users.data;
        this.storedData = users.data;
      } else {
        this.isLoading = false;
        this.isError = true;
      }
    });
  }

  changeFunctionCall() {
    this.frontendHandle = !this.frontendHandle;
  }

  handleInputChange(value) {
    if (value.length > 3) {
      this.inputValue = value;
      this.filteredLists = this.movies.filter((list) => {
        return list.name.toLowerCase().startsWith(value.toLowerCase());
      });
    } else this.filteredLists = this.movies;
  }

  handleInputChangeBackend(value) {
    if (value.length > 3) {
      this.isLoading = true;
      this.appService.getSpecificUser(value).subscribe((users: any) => {
        if ('data' in users) {
          this.isLoading = false;
          this.isDataAvailable = true;
          this.movies = users.data;
          this.filteredLists = users.data;
        } else {
          this.isLoading = false;
          this.isError = true;
        }
      });
    } else {
      this.isLoading = false;
      this.isDataAvailable = true;
      this.filteredLists = this.storedData;
    }
  }
}
