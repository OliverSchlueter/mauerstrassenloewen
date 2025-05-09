import { Routes } from '@angular/router';
import {RegisterComponent} from './register/register.component';
import {LoginComponent} from './login/login.component';
import {HomeComponent} from './home/home.component';
import {authGuard} from './auth.guard';
import {AccountComponent} from './home/account/account.component';
import {AccountLandingComponent} from './home/account/account-landing/account-landing.component';
import {AccountSettingsComponent} from './home/account/account-settings/account-settings.component';
import {AccountProfileComponent} from './home/account/account-profile/account-profile.component';
import {TheoryComponent} from './theory/theory.component';

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
    path: 'account',
    component: AccountComponent,
    children: [
      {
        path: 'landing',
        component: AccountLandingComponent,
      },
      {
        path: 'settings',
        component: AccountSettingsComponent
      },
      {
        path: 'profile',
        component: AccountProfileComponent
      }
    ]
  }
];
