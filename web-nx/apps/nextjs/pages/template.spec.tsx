import { render } from '@testing-library/react';

import TemplatePage from './template';

describe('TemplatePage', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<TemplatePage />);
    expect(baseElement).toBeTruthy();
  });
});
