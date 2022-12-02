import styles from './NextjsUtils.module.scss';

/* eslint-disable-next-line */
export interface NextjsUtilsProps {}

export function NextjsUtils(props: NextjsUtilsProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to NextjsUtils!</h1>
    </div>
  );
}

export default NextjsUtils;
