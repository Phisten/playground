import { Block, Marquee } from '@web-nx/ui';

export const Page = () => {
  return (
    <div className="grid">
      <Block title={'無限跑馬燈'}>
        {/** 須在tailwind.config.js內添加styles
          animation: {
            marquee: 'marquee 12s linear infinite; ',
          },
          keyframes: {
            marquee: {
              '0%': { transform: 'translateX(0%)' },
              '100%': { transform: 'translateX(-100%)' },
            },
          },
        */}
        <Marquee
          className="text-2xl"
          animationClass="animate-[marquee_8s_linear_infinite]"
        >
          {/* Marquee children 的兩倍寬度必須足夠填滿可能服務的螢幕寬度 */}
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
          {`Visualize Your Thoughts`}&nbsp;&nbsp;
        </Marquee>
      </Block>
    </div>
  );
};

export default Page;
