import { Component } from '@angular/core';
import {MatStep, MatStepLabel, MatStepper, MatStepperNext, MatStepperPrevious} from '@angular/material/stepper';
import {MatButton} from '@angular/material/button';
import {MatCard} from '@angular/material/card';

@Component({
  selector: 'app-register',
  imports: [
    MatStepper,
    MatStep,
    MatStepLabel,
    MatButton,
    MatStepperNext,
    MatStepperPrevious,
    MatCard
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {

  register() {

  }
}
