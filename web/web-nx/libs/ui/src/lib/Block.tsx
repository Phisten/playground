import clsx from 'clsx';
import { omit } from 'ramda';
import {
  DetailedHTMLProps,
  HTMLAttributes,
  PropsWithChildren,
  createContext,
  useContext,
} from 'react';

type DivProps = DetailedHTMLProps<
  HTMLAttributes<HTMLDivElement>,
  HTMLDivElement
>;

enum BlockConfigNames {
  default = 'default',
}
type BlockConfigName = keyof typeof BlockConfigNames | BlockConfigNames;
type BlockConfig = {
  [key in BlockConfigName]: Partial<BlockProps>;
};

const blockConfigContext = createContext<BlockConfig>({ default: {} });

export const useBlockConfig = (
  configName = 'default' as BlockConfigName,
  omitConfigProps?: Array<keyof BlockProps>
) => {
  const cfg = useContext(blockConfigContext)?.[configName];
  const newCfg: Partial<BlockProps> = omitConfigProps
    ? omit(omitConfigProps, cfg)
    : cfg;
  return newCfg;
};

export const BlockConfigProvider = (
  props: PropsWithChildren<{
    config: BlockConfig;
  }>
) => {
  const { children, config } = props;
  return (
    <blockConfigContext.Provider value={config}>
      {children}
    </blockConfigContext.Provider>
  );
};

/* eslint-disable-next-line */
export type BlockProps = PropsWithChildren<
  DivProps & {
    // spec
    title: string;
    titleClass?: string;
    //
    configName?: BlockConfigName;
    omitConfigProps?: Array<keyof BlockProps>;
  }
>;

export function Block(props: BlockProps) {
  const {
    // spec
    title,
    children,
    titleClass,
    // config
    configName,
    omitConfigProps,
    // general
    className,
    ...other
  } = props;
  const { className: configClass, ...otherConfig } = useBlockConfig(
    configName,
    omitConfigProps
  );

  return (
    <div className={clsx(configClass, className)} {...otherConfig} {...other}>
      {title ? <p className={titleClass}>{title}</p> : null}
      {children}
    </div>
  );
}

export default Block;
