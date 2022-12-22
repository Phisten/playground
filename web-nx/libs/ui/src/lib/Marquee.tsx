import clsx from 'clsx';
import { HTMLAttributes } from 'react';

/* eslint-disable-next-line */
export interface MarqueeProps extends HTMLAttributes<HTMLDivElement> {
  animationClass?: string;
}
/**
  theme: {
    extend: {
      animation: {
        marquee: 'marquee 12s linear infinite; ',
      },
      keyframes: {
        marquee: {
          '0%': { transform: 'translateX(0%)' },
          '100%': { transform: 'translateX(-100%)' },
        },
      },
    },
  },

  Marquee children 的兩倍長度必須足夠填滿可能服務的螢幕寬度 */
export function Marquee(props: MarqueeProps) {
  const { className, children, animationClass } = props;
  return (
    <div
      className={clsx(
        'flex justify-start overflow-hidden whitespace-nowrap',
        className
      )}
    >
      <div className={clsx('flex', animationClass)}>
        <span className="relative">
          {children}
          <span className="absolute left-full">{children}</span>
        </span>
      </div>
    </div>
  );
}

export default Marquee;
