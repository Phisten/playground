/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  createContext,
  Dispatch,
  PropsWithChildren,
  SetStateAction,
  useContext,
  useMemo,
  useState,
} from 'react';

type CartContextType = {
  records: Array<{ name: string }>;
  stepIndex: number;
};
type CartMutateType = {
  setRecord: Dispatch<CartContextType['records']>;
  setStepper: Dispatch<SetStateAction<number>>;

  create: (variables: any) => Promise<any>;
  clear: () => Promise<void>;
};
const CartContext = createContext<CartContextType>(
  null as unknown as CartContextType
);
const CartMutate = createContext<CartMutateType>(
  null as unknown as CartMutateType
);

export const useCartMutate = () => useContext<CartMutateType>(CartMutate);
export const useCartContext = () => ({
  ...useContext<CartContextType>(CartContext),
  ...useCartMutate(),
});

export const CartContextProvider = ({ children }: PropsWithChildren) => {
  const [records, setRecord] = useState<CartContextType['records']>([]);
  const [stepper, setStepper] = useState(0);

  const cartMutateValue: CartMutateType = useMemo(
    () => ({
      setRecord,
      setStepper,
      // eslint-disable-next-line @typescript-eslint/no-empty-function
      create: () => new Promise(() => {}),
      clear: () => new Promise(() => setRecord([])),
    }),
    []
  );

  return (
    <CartContext.Provider value={{ records, stepIndex: stepper }}>
      <CartMutate.Provider value={cartMutateValue}>
        {children}
      </CartMutate.Provider>
    </CartContext.Provider>
  );
};
