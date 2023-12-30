import { TableColumn } from '@mezzanine-ui/core/table';
import { Table } from '@mezzanine-ui/react';
import { useIsClient } from 'usehooks-ts';
import styles from './mzn.module.scss';
import clsx from 'clsx';

type DataType = typeof dataList[0];
const dataList = [
  { id: '0', name: 'a1', age: 12 },
  { id: '2', name: 'a12', age: 13123 },
  { id: '3', name: 'bb2', age: 22 },
  { id: '4', name: 'a3332', age: 6 },
  { id: '5', name: 'a144242', age: 20 },
  { id: '6', name: 'a2312', age: 19 },
];

const useColumns = (): TableColumn<DataType>[] => {
  return [
    {
      dataIndex: '',
      render(source) {
        return (
          <>
            <p className="z-[1]">{source.name}</p>
            {source.age > 20 ? (
              <div className="absolute bottom-0 left-0 h-[51px] w-full bg-red-400" />
            ) : null}
          </>
        );
      },
    },
    {
      dataIndex: 'age',
    },
  ];
};

export const Page = () => {
  const isClient = useIsClient();
  const columns = useColumns();

  return (
    <>
      <p>mznTableRow highlight</p>
      <a
        href={
          'https://codesandbox.io/s/mzn-table-row-highlight-j76m42?file=/src/App.js'
        }
        target={'_blank'}
        rel="noreferrer"
      >
        <p>codeSandbox</p>
      </a>
      <div className={clsx('w-screen h-screen', styles['wrapper'])}>
        {isClient && <Table dataSource={dataList} columns={columns} />}
      </div>
    </>
  );
};

export default Page;
