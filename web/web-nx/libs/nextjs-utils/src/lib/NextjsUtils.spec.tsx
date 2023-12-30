import { render } from '@testing-library/react';

import NextjsUtils from './NextjsUtils';

describe('NextjsUtils', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<NextjsUtils />);
    expect(baseElement).toBeTruthy();
  });
});
