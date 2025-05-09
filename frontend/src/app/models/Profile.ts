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
  primaryGoal: 'Wealth Building' | 'Side Income' | 'Financial Freedom' | 'Retirement Planning' | 'Speculation' | '' = '';
  timeHorizon: 'Short-Term' | 'Medium-Term' | 'Long-Term' | '' = '';
  startingCapital: string = '';
  monthlyInvestment: string = '';
  annualProfitTarget: string = '';

  // Trading Preferences
  preferredMarkets: ('Stocks' | 'ETFs' | 'Crypto' | 'Forex' | 'Commodities' | 'Options' | 'Futures' | 'CFDs')[] = [];
  toolsPlatforms: string = '';
  weeklyTimeCommitment: string = '';

  // Experience and Knowledge
  tradingExperience: 'Beginner' | 'Intermediate' | 'Expert' | '' = '';
  pastPerformance: string = '';
  chartAnalysisKnowledge: 'None' | 'Basic' | 'Advanced' | 'Expert' | '' = '';
  fundamentalAnalysisKnowledge: 'None' | 'Basic' | 'Advanced' | 'Expert' | '' = '';
  certifications: string = '';

  // Psychological Profile
  riskType: 'Risk-Taking' | 'Balanced' | 'Cautious' | '' = '';
  lossBehavior: string = '';
  discipline: string = '';
  tradingJournal: boolean | null = null;
  stressTolerance: 'High' | 'Medium' | 'Low' | '' = '';

  // Ethics and Values
  esgPreferences: 'Important' | 'Neutral' | 'Unimportant' | '' = '';
  exclusionCriteria: string = '';
  politicalInvesting: boolean | null = null;
  religiousRestrictions: string = '';

  // Technical Setup
  bookkeeping: 'Excel' | 'Trading Journal Software' | 'None' | '' = '';
  availableTech: string = '';
}

