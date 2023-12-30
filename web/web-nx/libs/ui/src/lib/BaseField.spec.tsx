import { render } from '@testing-library/react';

import BaseField from './BaseField';

describe('BaseField', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<BaseField />);
    expect(baseElement).toBeTruthy();
  });
});
