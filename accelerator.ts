interface Accelerator {
  name: string;
  equity: string;
  website: string;
  location: string;
  focusAreas: string[];
  investmentRange: string;
  description: string;
}

const europeanEdTechAccelerators: Accelerator[] = [
  // European accelerators and incubators focused on EdTech and language learning
  // Data compiled from comprehensive research as of 2024
];

export { Accelerator, europeanEdTechAccelerators };