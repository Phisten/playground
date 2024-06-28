// 紅色 grid-column 的寬度 =  [螢幕寬-左右 Margin - ((Count - 1) x Gutter 寬)]/Count

import { useWindowSize } from 'usehooks-ts';
import useIsHydrated from './useIsHydrated';

/**
--grid-system-margin: 整個 grid-system 外部的空白處寬度(左右 margin or padding);
--grid-system-column-width: width100%包含margin時,每條column的寬度;
--grid-system-inner-column-width: width100%不包含margin時,每條column的寬度;
--grid-system-gap: columns 之間的間格(gutter寬度);
--grid-system-width: grid-system 所有column + (count-1)條 gutter 的總寬度;
--grid-system-max-width: grid-system 的最大寬度(螢幕寬度大於一定值時 columns 不再放大或增加數量);
*/
export const gridSystem = {
  /* grid: column count */
  $column_count_max_960: 12,
  $column_count_959_568: 8,
  $column_count_567_0: 6,
  /* grid: max width */
  $content_max_width: 1248, // 1680以上固定寬
  /* grid: column gutter */
  $gutter_max_960: 40,
  $gutter_959_568: 20, // 設計稿上沒標注 需確認
  $gutter_567_0: 16, // 設計稿上沒標注 需確認
  /* grid: margin */
  $margin_1679_1440: 64,
  $margin_1439_568: 48,
  $margin_576_0: 20,

  /* RWD breakpoints */
  $pc1680: 1680,
  $pc1440: 1440,
  $pc1280: 1280,
  // $pc1024: 1024,
  $pad960: 960,
  $pad768: 768,
  $pad568: 568,
  $mob0: 0,
};

export const useGridSystem = () => {
  const isHydrated = useIsHydrated();
  const { width: windowWidth, height: windowHeight } = useWindowSize();

  let gsMargin = (windowWidth - gridSystem.$content_max_width) / 2;
  let gsGap = gridSystem.$gutter_max_960;
  let gsColumnCount = gridSystem.$column_count_max_960;

  const widthInRange = (under: number, on: number) => windowWidth < under && windowWidth >= on;

  if (widthInRange(gridSystem.$pc1680, gridSystem.$pc1440)) {
    gsMargin = gridSystem.$margin_1679_1440;
  }
  if (widthInRange(gridSystem.$pc1440, gridSystem.$pad960)) {
    gsMargin = gridSystem.$margin_1439_568;
  }
  if (widthInRange(gridSystem.$pad960, gridSystem.$pad568)) {
    gsMargin = gridSystem.$margin_1439_568;
    gsGap = gridSystem.$gutter_959_568;
    gsColumnCount = gridSystem.$column_count_959_568;
  }
  if (widthInRange(gridSystem.$pad568, gridSystem.$mob0)) {
    gsMargin = gridSystem.$margin_576_0;
    gsGap = gridSystem.$gutter_567_0;
    gsColumnCount = gridSystem.$column_count_567_0;
  }

  const gsWidth = (windowWidth || 375) - gsMargin - gsMargin;
  // 100%包含margin時,每條column的寬度
  const gsColumnWidth = (gsWidth - (gsColumnCount - 1) * gsGap) / gsColumnCount;
  // 100%不包含margin時,每條column的寬度

  const calcGsInnerColumnWidth = (containerWidth: number) =>
    (containerWidth - (gsColumnCount - 1) * gsGap) / gsColumnCount;

  return {
    isHydrated,
    windowWidth,
    windowHeight,
    gsWidth,
    gsMargin,
    gsGap,
    gsColumnCount,
    gsColumnWidth,
    calcGsInnerColumnWidth,
  };
};
