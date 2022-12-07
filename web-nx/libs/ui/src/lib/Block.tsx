import { DetailedHTMLProps, HTMLAttributes, PropsWithChildren } from 'react';

type DivProps = DetailedHTMLProps<
  HTMLAttributes<HTMLDivElement>,
  HTMLDivElement
>;

/* eslint-disable-next-line */
export type BlockProps = PropsWithChildren<
  DivProps & { title: string; titleClass?: string }
>;

export function Block(props: BlockProps) {
  const { title, children, titleClass, ...other } = props;
  return (
    <div {...other}>
      {title ? <p className={titleClass}>{title}</p> : null}
      {children}
    </div>
  );
}

export default Block;
