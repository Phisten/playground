import { rm } from 'fs/promises';
import { join } from 'path';

global.beforeEach(async () => {
  try {
    console.log('delete db-test.sqlite');
    await rm(join(__dirname, '..', 'db-test.sqlite'));
  } catch (error) {}
});
