package script

import (
    "fmt"
    "strconv"
    "strings"
    "unicode/utf16"
)

func GetEmoji(texts string) string {
    group := strings.Split(texts, `\u`)

    if len(group) == 3 {
        n0 := toUint16(group[1])
        n1 := toUint16(group[2])

        emoji := utf16.Decode([]uint16{n0, n1})
        emojiStr := fmt.Sprintf(`%x`, emoji)

        emojiStr = Scan(emojiStr, `(\w+)`, 1)

        return emojiMap[emojiStr]
    }
    return ""
}

func toUint16(str string) uint16 {
    n, _ := strconv.ParseInt(str, 16, 32)
    return uint16(n)
}

var (
    emojiMap = map[string]string{
        "1f440": "👀",
        "1f441": "👁",
        "1f442": "👂",
        "1f443": "👃",
        "1f444": "👄",
        "1f445": "👅",
        "1f446": "👆",
        "1f447": "👇",
        "1f448": "👈",
        "1f449": "👉",
        "1f44a": "👊",
        "1f44b": "👋",
        "1f44c": "👌",
        "1f44d": "👍",
        "1f44e": "👎",
        "1f44f": "👏",
        "1f450": "👐",
        "1f479": "👹",
        "1f47a": "👺",
        "1f47b": "👻",
        "1f47d": "👽",
        "1f47e": "👾",
        "1f47f": "👿",
        "1f480": "💀",
        "1f485": "💅",
        "1f48b": "💋",
        "1f48c": "💌",
        "1f493": "💓",
        "1f494": "💔",
        "1f495": "💕",
        "1f496": "💖",
        "1f497": "💗",
        "1f498": "💘",
        "1f499": "💙",
        "1f49a": "💚",
        "1f49b": "💛",
        "1f49c": "💜",
        "1f49d": "💝",
        "1f49e": "💞",
        "1f49f": "💟",
        "1f4a2": "💢",
        "1f4a4": "💤",
        "1f4a5": "💥",
        "1f4a6": "💦",
        "1f4a8": "💨",
        "1f4a9": "💩",
        "1f4aa": "💪",
        "1f4ab": "💫",
        "1f4ac": "💬",
        "1f4ad": "💭",
        "1f4af": "💯",
        "1f573": "🕳",
        "1f590": "🖐",
        "1f595": "🖕",
        "1f596": "🖖",
        "1f5a4": "🖤",
        "1f5e8": "🗨",
        "1f5ef": "🗯",
        "1f600": "😀",
        "1f601": "😁",
        "1f602": "😂",
        "1f603": "😃",
        "1f604": "😄",
        "1f605": "😅",
        "1f606": "😆",
        "1f607": "😇",
        "1f608": "😈",
        "1f609": "😉",
        "1f60a": "😊",
        "1f60b": "😋",
        "1f60c": "😌",
        "1f60d": "😍",
        "1f60e": "😎",
        "1f60f": "😏",
        "1f610": "😐",
        "1f611": "😑",
        "1f612": "😒",
        "1f613": "😓",
        "1f614": "😔",
        "1f615": "😕",
        "1f616": "😖",
        "1f617": "😗",
        "1f618": "😘",
        "1f619": "😙",
        "1f61a": "😚",
        "1f61b": "😛",
        "1f61c": "😜",
        "1f61d": "😝",
        "1f61e": "😞",
        "1f61f": "😟",
        "1f620": "😠",
        "1f621": "😡",
        "1f622": "😢",
        "1f623": "😣",
        "1f624": "😤",
        "1f625": "😥",
        "1f626": "😦",
        "1f627": "😧",
        "1f628": "😨",
        "1f629": "😩",
        "1f62a": "😪",
        "1f62b": "😫",
        "1f62c": "😬",
        "1f62d": "😭",
        "1f62e": "😮",
        "1f62f": "😯",
        "1f630": "😰",
        "1f631": "😱",
        "1f632": "😲",
        "1f633": "😳",
        "1f634": "😴",
        "1f635": "😵",
        "1f636": "😶",
        "1f637": "😷",
        "1f638": "😸",
        "1f639": "😹",
        "1f63a": "😺",
        "1f63b": "😻",
        "1f63c": "😼",
        "1f63d": "😽",
        "1f63e": "😾",
        "1f63f": "😿",
        "1f640": "🙀",
        "1f641": "🙁",
        "1f642": "🙂",
        "1f643": "🙃",
        "1f644": "🙄",
        "1f648": "🙈",
        "1f649": "🙉",
        "1f64a": "🙊",
        "1f64c": "🙌",
        "1f64f": "🙏",
        "1f90c": "🤌",
        "1f90d": "🤍",
        "1f90e": "🤎",
        "1f90f": "🤏",
        "1f910": "🤐",
        "1f911": "🤑",
        "1f912": "🤒",
        "1f913": "🤓",
        "1f914": "🤔",
        "1f915": "🤕",
        "1f916": "🤖",
        "1f917": "🤗",
        "1f918": "🤘",
        "1f919": "🤙",
        "1f91a": "🤚",
        "1f91b": "🤛",
        "1f91c": "🤜",
        "1f91d": "🤝",
        "1f91e": "🤞",
        "1f91f": "🤟",
        "1f920": "🤠",
        "1f921": "🤡",
        "1f922": "🤢",
        "1f923": "🤣",
        "1f924": "🤤",
        "1f925": "🤥",
        "1f927": "🤧",
        "1f928": "🤨",
        "1f929": "🤩",
        "1f92a": "🤪",
        "1f92b": "🤫",
        "1f92c": "🤬",
        "1f92d": "🤭",
        "1f92e": "🤮",
        "1f92f": "🤯",
        "1f932": "🤲",
        "1f933": "🤳",
        "1f970": "🥰",
        "1f971": "🥱",
        "1f972": "🥲",
        "1f973": "🥳",
        "1f974": "🥴",
        "1f975": "🥵",
        "1f976": "🥶",
        "1f978": "🥸",
        "1f979": "🥹",
        "1f97a": "🥺",
        "1f9b4": "🦴",
        "1f9b5": "🦵",
        "1f9b6": "🦶",
        "1f9b7": "🦷",
        "1f9bb": "🦻",
        "1f9be": "🦾",
        "1f9bf": "🦿",
        "1f9d0": "🧐",
        "1f9e0": "🧠",
        "1f9e1": "🧡",
        "1fa75": "🩵",
        "1fa76": "🩶",
        "1fa77": "🩷",
        "1fac0": "🫀",
        "1fac1": "🫁",
        "1fae0": "🫠",
        "1fae1": "🫡",
        "1fae2": "🫢",
        "1fae3": "🫣",
        "1fae4": "🫤",
        "1fae5": "🫥",
        "1fae6": "🫦",
        "1fae8": "🫨",
        "1faf0": "🫰",
        "1faf1": "🫱",
        "1faf2": "🫲",
        "1faf3": "🫳",
        "1faf4": "🫴",
        "1faf5": "🫵",
        "1faf6": "🫶",
        "1faf7": "🫷",
        "1faf8": "🫸",
        "261d":  "☝",
        "2620":  "☠",
        "2639":  "☹",
        "263a":  "☺",
        "270a":  "✊",
        "270b":  "✋",
        "270c":  "✌",
        "270d":  "✍",
        "2763":  "❣",
        "2764":  "❤",
    }
)
