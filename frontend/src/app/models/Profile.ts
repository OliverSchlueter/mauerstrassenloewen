export class Profile {
  // Personal Information
  id: string = '';
  surname: string = '';
  name: string = '';
  birthdate: Date | null = null;
  profession: string = '';
  employmentStatus: 'Employed' | 'Self-Employed' | 'Entrepreneur' | 'Job-Seeking' | 'Student' | 'Retired' | '' = '';
  country: string = '';
  maritalStatus: 'Single' | 'Married' | 'With Children' | 'Single Parent' | 'Other' | '' = '';
  financialObligations: string = '';

  // Goals and Objectives
  primaryGoal: 'Wealth Building' | 'Side Income' | 'Financial Freedom' | 'Retirement Planning' | '' = '';
  timeHorizon: 'Short-Term' | 'Medium-Term' | 'Long-Term' | '' = '';
  startingCapital: string = '';
  monthlyInvestment: string = '';

  // Commitment and Experience
  preferredMarkets: ('Stocks' | 'ETFs' | 'Crypto' | 'Forex' | 'Commodities' | 'Options' | 'Futures' | 'CFDs')[] = [];
  toolsPlatforms: string = '';
  weeklyTimeCommitment: string = '';

  tradingExperience: 'None' | 'Beginner' | 'Intermediate' | 'Expert' | '' = '';
  chartAnalysisKnowledge: 'None' | 'Basic' | 'Advanced' | 'Expert' | '' = '';
  fundamentalAnalysisKnowledge: 'None' | 'Basic' | 'Advanced' | 'Expert' | '' = '';
  certifications: string = '';

  // Psychological Profile
  riskAffinity: number | null = null;
  lossBehavior: string = '';
  stressTolerance: 'High' | 'Medium' | 'Low' | '' = '';

  // Ethics and Values
  esgPreferences: 'Important' | 'Medium' | 'Unimportant' | '' = '';
  exclusionCriteria: string = '';
  excludePoliticalInvesting: boolean | null = null;
  religiousRestrictions: string = '';

  // Technical Setup
  tradingJournal: boolean | null = null;
  informationManagement: 'Excel' | 'Trading Journal Software' | 'None' | '' = '';
  availableTech: string = '';
}

