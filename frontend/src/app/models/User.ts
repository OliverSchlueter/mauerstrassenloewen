import {Profile} from './Profile';

export class User {
  id: string = '';
  name: string = '';
  email: string = '';
  password: string = '';
  profile: Profile = new Profile();
}
