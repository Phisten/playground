import Link, { LinkProps } from 'next/link';
import { ReactNode } from 'react';

export const NextLink = (props: LinkProps & { children?: ReactNode }) => {
  return (
    <Link {...props} passHref>
      <a href="passHref">{props?.children}</a>
    </Link>
  );
};
