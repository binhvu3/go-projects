import { background, extendTheme, ThemeConfig } from "@chakra-ui/react";
import { mode } from "@chakra-ui/theme-tools";

const config: ThemeConfig = {
  initialColorMode: "dark",
  useSystemColorMode: true,
};

const theme = extendTheme({
  config,
  styles: {
    global: (props:any) => ({
        body:{
            backgroundColor: mode("gray.400", "")(props),
        }
    })
  }
});

export default theme;
