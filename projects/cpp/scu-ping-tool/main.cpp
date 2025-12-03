#define _WINSOCK_DEPRECATED_NO_WARNINGS

// 标准库头文件
#include <algorithm>
#include <chrono>
#include <cstdint>
#include <cstring>
#include <iomanip>
#include <iostream>
#include <string>
#include <vector>

// Windows 系统头文件
#include <winsock2.h>
#include <ws2tcpip.h>

// clang-format off
// IP Helper 和 ICMP API 头文件
#include <iphlpapi.h>
#include <icmpapi.h>
// clang-format on

// 链接 Winsock 和 IP Helper 库
#pragma comment(lib, "ws2_32.lib")
#pragma comment(lib, "iphlpapi.lib")

// ICMP 报头结构定义
struct IcmpHeader {
  uint8_t type;       // 类型 (8 = 回显请求 Echo Request)
  uint8_t code;       // 代码 (0)
  uint16_t checksum;  // 校验和
  uint16_t id;        // 标识符
  uint16_t seq;       // 序列号
};

// 计算缓冲区的 Internet 校验和 (大小以字节为单位)
static uint16_t CalculateChecksum(const uint8_t* buffer, int size) {
  uint32_t sum          = 0;
  const uint16_t* words = reinterpret_cast<const uint16_t*>(buffer);

  while (size > 1) {
    sum += *words++;
    size -= 2;
  }

  if (size == 1) {
    // 处理末尾的奇数字节
    const uint8_t last_byte = *(reinterpret_cast<const uint8_t*>(words));
    sum += static_cast<uint16_t>(last_byte) << 8;
  }

  // 将 32 位和折叠为 16 位并返回反码
  sum = (sum >> 16) + (sum & 0xFFFF);
  sum += (sum >> 16);
  return static_cast<uint16_t>(~sum);
}

// 打印用法说明
static void PrintUsage() {
  std::cout << "用法: scu-ping-tool target_name [-t] [-n count] [-l size]\n";
  std::cout << "选项:\n";
  std::cout << "    -t          Ping 指定的主机，直到停止。\n";
  std::cout << "    -n count    要发送的回显请求数。\n";
  std::cout << "    -l size     发送缓冲区大小。\n";
}

int main(int argc, char** argv) {
  std::string target;
  int ping_count     = 4;   // 默认发送 4 次
  int data_size      = 32;  // 默认数据大小 32 字节
  bool infinite_ping = false;

  // 解析命令行参数
  if (argc < 2) {
    std::cout << "请输入要 ping 的目标主机 (例如: www.baidu.com): ";
    std::cin >> target;
  } else {
    target = argv[1];
    for (int i = 2; i < argc; ++i) {
      std::string arg = argv[i];
      if (arg == "-t") {
        infinite_ping = true;
      } else if (arg == "-n" && i + 1 < argc) {
        ping_count = std::stoi(argv[++i]);
      } else if (arg == "-l" && i + 1 < argc) {
        data_size = std::stoi(argv[++i]);
      } else if (arg == "-h" || arg == "--help") {
        PrintUsage();
        return 0;
      }
    }
  }

  // 初始化 Winsock
  WSADATA wsa_data;
  if (WSAStartup(MAKEWORD(2, 2), &wsa_data) != 0) {
    std::cerr << "WSAStartup 失败。\n";
    return 1;
  }

  // 创建原始套接字 (Raw Socket)
  SOCKET sock = socket(AF_INET, SOCK_RAW, IPPROTO_ICMP);
  if (sock == INVALID_SOCKET) {
    std::cerr << "创建套接字失败: " << WSAGetLastError() << "\n";
    std::cerr << "注意: 原始套接字通常需要管理员权限运行。\n";
    WSACleanup();
    return 1;
  }

  // 设置接收超时时间 (毫秒)
  int timeout_ms = 1000;
  setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO,
             reinterpret_cast<const char*>(&timeout_ms), sizeof(timeout_ms));

  // 准备目标地址结构
  sockaddr_in dest_addr;
  std::memset(&dest_addr, 0, sizeof(dest_addr));
  dest_addr.sin_family = AF_INET;

  // 解析目标 IP 地址
  dest_addr.sin_addr.s_addr = inet_addr(target.c_str());
  if (dest_addr.sin_addr.s_addr == INADDR_NONE) {
    hostent* host = gethostbyname(target.c_str());
    if (host) {
      dest_addr.sin_addr = *reinterpret_cast<in_addr*>(host->h_addr);
    } else {
      std::cerr << "Ping 请求找不到主机 " << target
                << "。请检查该名称，然后重试。\n";
      closesocket(sock);
      WSACleanup();
      return 1;
    }
  }

  std::cout << "\n正在 Ping " << target << " [" << inet_ntoa(dest_addr.sin_addr)
            << "] 具有 " << data_size << " 字节的数据:\n\n";

  const int packet_size = static_cast<int>(sizeof(IcmpHeader)) + data_size;
  std::vector<uint8_t> send_buf(packet_size);
  // 接收缓冲区需要包含 IP 报头的空间
  std::vector<uint8_t> recv_buf(packet_size + 100);

  int sent_count          = 0;
  int recv_count          = 0;
  long long min_time_us   = LLONG_MAX;
  long long max_time_us   = 0;
  long long total_time_us = 0;

  int seq                   = 0;
  const uint16_t identifier = static_cast<uint16_t>(GetCurrentProcessId());

  // 主 Ping 循环
  while (infinite_ping || seq < ping_count) {
    std::fill(send_buf.begin(), send_buf.end(), 0);

    // 填充 ICMP 报头
    IcmpHeader* icmp = reinterpret_cast<IcmpHeader*>(send_buf.data());
    icmp->type       = 8;  // 回显请求 (Echo Request)
    icmp->code       = 0;
    icmp->id         = identifier;
    icmp->seq        = static_cast<uint16_t>(seq);
    icmp->checksum   = 0;

    // 填充数据部分
    std::fill(send_buf.begin() + sizeof(IcmpHeader), send_buf.end(),
              static_cast<uint8_t>('E'));

    // 计算校验和
    icmp->checksum = CalculateChecksum(send_buf.data(), packet_size);

    auto start = std::chrono::high_resolution_clock::now();

    // 发送数据包
    int sent_bytes = sendto(
        sock, reinterpret_cast<const char*>(send_buf.data()), packet_size, 0,
        reinterpret_cast<sockaddr*>(&dest_addr), sizeof(dest_addr));
    if (sent_bytes == SOCKET_ERROR) {
      std::cerr << "发送失败: " << WSAGetLastError() << "\n";
    } else {
      ++sent_count;
    }

    // 接收响应
    sockaddr_in from_addr;
    int from_len = sizeof(from_addr);
    int recv_bytes =
        recvfrom(sock, reinterpret_cast<char*>(recv_buf.data()),
                 static_cast<int>(recv_buf.size()), 0,
                 reinterpret_cast<sockaddr*>(&from_addr), &from_len);

    auto end = std::chrono::high_resolution_clock::now();
    long long duration_us =
        std::chrono::duration_cast<std::chrono::microseconds>(end - start)
            .count();

    if (recv_bytes == SOCKET_ERROR) {
      if (WSAGetLastError() == WSAETIMEDOUT) {
        std::cout << "请求超时。\n";
      } else {
        std::cout << "一般故障。\n";
      }
    } else {
      // 解析接收到的数据包
      // IP 报头长度位于第一个字节的低 4 位 (单位为 4 字节)
      int ip_header_len =
          (static_cast<int>(static_cast<unsigned char>(recv_buf[0])) & 0x0f) *
          4;
      if (recv_bytes >= ip_header_len + static_cast<int>(sizeof(IcmpHeader))) {
        IcmpHeader* recv_icmp =
            reinterpret_cast<IcmpHeader*>(recv_buf.data() + ip_header_len);

        // 检查是否为回显应答 (Type 0) 且 ID 匹配
        if (recv_icmp->type == 0 && recv_icmp->id == icmp->id) {
          ++recv_count;
          total_time_us += duration_us;
          min_time_us = (std::min)(min_time_us, duration_us);
          max_time_us = (std::max)(max_time_us, duration_us);

          double duration_ms = duration_us / 1000.0;
          std::cout << "来自 " << inet_ntoa(from_addr.sin_addr) << " 的回复: "
                    << "字节=" << data_size << " "
                    << "时间="
                    << (duration_ms < 1.0
                            ? "<1"
                            : std::to_string(static_cast<int>(duration_ms)))
                    << "ms "
                    << "TTL="
                    << static_cast<int>(static_cast<unsigned char>(recv_buf[8]))
                    << "\n";
        } else {
          std::cout << "收到非预期的数据包。\n";
        }
      } else {
        std::cout << "收到的数据包长度不足，无法解析。\n";
      }
    }

    ++seq;
    if (infinite_ping || seq < ping_count) {
      Sleep(1000);
    }
  }

  // 打印统计信息
  std::cout << "\n" << target << " 的 Ping 统计信息:\n";
  std::cout << "    数据包: 已发送 = " << sent_count
            << "，已接收 = " << recv_count
            << "，丢失 = " << (sent_count - recv_count) << " ("
            << (sent_count > 0
                    ? static_cast<int>(
                          (static_cast<double>(sent_count - recv_count) /
                           sent_count) *
                          100)
                    : 0)
            << "% 丢失)，\n";

  if (recv_count > 0) {
    std::cout << "往返行程的估计时间(以毫秒为单位):\n";
    std::cout << std::fixed << std::setprecision(2);
    std::cout << "    最短 = " << (min_time_us / 1000.0) << "ms，"
              << "最长 = " << (max_time_us / 1000.0) << "ms，"
              << "平均 = " << (total_time_us / recv_count / 1000.0) << "ms\n";
  }

  closesocket(sock);
  WSACleanup();
  return 0;
}
