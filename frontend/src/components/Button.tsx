import React, { ComponentProps } from 'react';
import cn from 'classnames';

export interface ButtonProps extends ComponentProps<'button'> {
  buttonStyleType?: ButtonStypeType;
}

export const Button: React.FunctionComponent<ButtonProps> = (props) => {
  const extended: ComponentProps<'button'> = {
    ...props,
    className: cn(getButtonClasses(props.buttonStyleType), props.className),
  };
  return <button {...extended}>{props.children}</button>;
};

export function getButtonClasses(style: ButtonStypeType = 'primary'): string {
  switch (style) {
    case 'outline':
      return 'px-5 py-3 rounded-2xl border border-primary text-primary';
    case 'primary':
      return 'px-5 py-3 rounded-2xl bg-primary text-white';
  }
}

export type ButtonStypeType = 'primary' | 'outline';
