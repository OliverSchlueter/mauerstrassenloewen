import { Routes } from '@angular/router';
import {RegisterComponent} from './register/register.component';
import {LoginComponent} from './login/login.component';
import {HomeComponent} from './home/home.component';
import {authGuard} from './auth.guard';
import {AccountComponent} from './home/account/account.component';
import {TheoryComponent} from './theory/theory.component';
import {InterviewComponent} from './interview/interview.component';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full'
  },
  {
    path: 'home',
    component: HomeComponent,
    canActivate: [authGuard]
  },
  {
    path: 'theory',
    component: TheoryComponent,
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'register',
    component: RegisterComponent,
  },
  {
    path: 'interview',
    component: InterviewComponent,
  },
  {
    path: 'account',
    component: AccountComponent,
  }
];
