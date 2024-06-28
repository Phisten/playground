import { DependencyList, EffectCallback, useEffect, useState } from 'react';

/** 當環境變為瀏覽器時強制重新渲染，回傳值為目前是否已可渲染伺服器無法渲染的元件(依賴瀏覽器變數window的元件)。初次調用必定為false(防止伺服端與用戶端渲染的結果相異導致的hydrate問題)
 *
 * @more-info-react https://reactjs.org/docs/react-dom.html#hydrate
 * @more-info-NextJs https://nextjs.org/docs/messages/react-hydration-error
 *
 * @more-info-ssr-useLayoutEffect https://gist.github.com/gaearon/e7d97cdf38a2907924ea12e4ebdf3c85
 *
 * @note 類似的解決方案 https://usehooks-ts.com/  react-hook/use-is-client
 */
export const useIsHydrated = (effect?: EffectCallback, deps?: DependencyList) => {
  const [isHydrated, checkHydratedRerender] = useState(false);

  useEffect(() => {
    checkHydratedRerender(true);
  }, []);

  return isHydrated;
};

/** 水合完成後再次重渲染, 可用於設定依賴於瀏覽器的初始值 */
export const useHydratedEffect = (effect: EffectCallback, deps?: DependencyList) => {
  const isHydrated = useIsHydrated();

  useEffect(() => {
    if (isHydrated) {
      effect();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isHydrated, ...(deps ?? [])]);

  return isHydrated;
};

export default useIsHydrated;
