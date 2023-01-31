import clsx from 'clsx';
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

const configNames = {
  default: 'default',
};
type configNameType = keyof typeof configNames;
type configType = {
  [key in configNameType]: Partial<BlockProps>;
};

const blockConfigContext = createContext<configType>({ default: {} });

export const useBlockConfig = (
  configName = 'default' as configNameType,
  omitConfigProps?: Array<keyof BlockProps>
) => {
  const cfg = useContext(blockConfigContext)?.[configName];
  omitConfigProps?.forEach((p) => (cfg[p] = undefined));

  return cfg;
};

export const BlockConfigProvider = (
  props: PropsWithChildren<{
    config: configType;
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
    title: string;
    titleClass?: string;
    configName?: configNameType;
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
